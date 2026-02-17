// --- THREE.JS 3D SCENE SETUP ---
const container = document.getElementById('xray-container');
const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, container.clientWidth / container.clientHeight, 0.1, 1000);
const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });

renderer.setSize(container.clientWidth, container.clientHeight);
container.appendChild(renderer.domElement);

// ПАРАМЕТРЫ ДИНАМИКИ
let targetRotationSpeed = 0.01;
let currentRotationSpeed = 0.01;
let coreOpacity = 0.8;

// Создаем "Sentinel Core" — октаэдр
const geometry = new THREE.OctahedronGeometry(1.5, 0);
const material = new THREE.MeshBasicMaterial({ 
    color: 0x00ff41, 
    wireframe: true,
    transparent: true,
    opacity: coreOpacity
});
const core = new THREE.Mesh(geometry, material);
scene.add(core);

// Внешняя оболочка
const outerGeo = new THREE.IcosahedronGeometry(2.2, 1);
const outerMat = new THREE.MeshBasicMaterial({ 
    color: 0x00ff41, 
    wireframe: true, 
    transparent: true, 
    opacity: 0.1 
});
const shield = new THREE.Mesh(outerGeo, outerMat);
scene.add(shield);

camera.position.z = 5;

// Функция для импульса (ускорения)
function triggerTransactionPulse() {
    targetRotationSpeed = 0.15; // Резко ускоряем
    material.opacity = 1.0;     // Делаем ярче
    document.getElementById('terminal-text').innerText = "TRANSACTION DETECTED: ACCELERATING FPGA...";
    
    // Через 800мс возвращаем к нормальной скорости
    setTimeout(() => {
        targetRotationSpeed = 0.01;
        material.opacity = 0.8;
    }, 800);
}

// Анимация
function animate3D() {
    requestAnimationFrame(animate3D);
    
    // Плавный переход скорости (интерполяция)
    currentRotationSpeed += (targetRotationSpeed - currentRotationSpeed) * 0.1;
    
    core.rotation.y += currentRotationSpeed;
    core.rotation.x += currentRotationSpeed / 2;
    shield.rotation.y -= 0.003;
    
    renderer.render(scene, camera);
}
animate3D();

// --- LIVE DATA SIMULATION ---

// 1. Энтропия
setInterval(() => {
    const display = document.getElementById('entropy-display');
    let bytes = "";
    for(let i=0; i<4; i++) {
        bytes += Math.random().toString(16).slice(2, 6).toUpperCase() + " ";
    }
    display.innerText = bytes;
}, 100);

// 2. Финальность + Случайные импульсы транзакций
setInterval(() => {
    const val = (Math.random() * (0.98 - 0.72) + 0.72).toFixed(2);
    document.getElementById('ms-value').innerText = val;
    document.getElementById('finality-bar').style.width = (val * 100) + "%";
    
    // С шансом 10% имитируем транзакцию для визуала
    if(Math.random() > 0.9) {
        triggerTransactionPulse();
    }
}, 150);

// 3. Терминал (Статусы)
const logs = [
    "EXECUTING POSEIDON HASH...",
    "FPGA TEMP: 31.4°C [OPTIMAL]",
    "L-BAND SIGNAL: STABLE",
    "ATOMIC SYNC: <1ps DRIFT",
    "SENTINEL LITE: PARALLEL OK"
];
let logIdx = 0;
setInterval(() => {
    if(targetRotationSpeed < 0.05) { // Не перебивать лог транзакции
        document.getElementById('terminal-text').innerText = logs[logIdx];
        logIdx = (logIdx + 1) % logs.length;
    }
}, 2500);

// 4. Panic Switch
document.getElementById('panic-switch').addEventListener('click', () => {
    const confirmed = confirm("WARNING: IRREVERSIBLE KEY DISSOLUTION. PROCEED?");
    if(confirmed) {
        targetRotationSpeed = 0.5; // Безумное вращение перед смертью
        document.body.style.filter = "invert(1) hue-rotate(180deg)";
        document.getElementById('terminal-text').innerText = "SYSTEM PURGED. SHUTTING DOWN...";
        setTimeout(() => location.reload(), 2000);
    }
});

// Resize handler
window.addEventListener('resize', () => {
    camera.aspect = container.clientWidth / container.clientHeight;
    camera.updateProjectionMatrix();
    renderer.setSize(container.clientWidth, container.clientHeight);
});
