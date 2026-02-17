// --- THREE.JS 3D SCENE SETUP ---
const container = document.getElementById('xray-container');
const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, container.clientWidth / container.clientHeight, 0.1, 1000);
const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });

renderer.setSize(container.clientWidth, container.clientHeight);
container.appendChild(renderer.domElement);

// Параметры динамики ядра
let targetRotationSpeed = 0.01;
let currentRotationSpeed = 0.01;

// Создаем "Sentinel Core" — октаэдр (внутренний чип)
const geometry = new THREE.OctahedronGeometry(1.5, 0);
const material = new THREE.MeshBasicMaterial({ 
    color: 0x00ff41, 
    wireframe: true,
    transparent: true,
    opacity: 0.8
});
const core = new THREE.Mesh(geometry, material);
scene.add(core);

// Внешняя оболочка (защитный слой)
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

// Функция импульса при транзакции
function triggerTransactionPulse() {
    targetRotationSpeed = 0.2; // Резкий рывок
    material.opacity = 1.0;
    document.getElementById('terminal-text').innerText = "TRANSACTION DETECTED: ACCELERATING SENTINEL CORE...";
    
    setTimeout(() => {
        targetRotationSpeed = 0.01;
        material.opacity = 0.8;
    }, 1000);
}

// Рендер-цикл
function animate() {
    requestAnimationFrame(animate);
    
    // Плавное изменение скорости (инерция)
    currentRotationSpeed += (targetRotationSpeed - currentRotationSpeed) * 0.1;
    
    core.rotation.y += currentRotationSpeed;
    core.rotation.x += currentRotationSpeed / 2;
    shield.rotation.y -= 0.002;
    
    renderer.render(scene, camera);
}
animate();

// --- СВЯЗЬ С ТЕСТНЕТОМ (WEBSOCKET) ---

// Замени 'localhost' на IP своего сервера, когда выйдешь в онлайн
const socket = new WebSocket('ws://localhost:8080/ws');

socket.onopen = () => {
    console.log("Connected to TOTAL Heartbeat Server");
    document.getElementById('terminal-text').innerText = "CONNECTION ESTABLISHED. CORE SYNCED.";
};

socket.onmessage = (event) => {
    const data = JSON.parse(event.data);

    // 1. Обновляем энтропию из сервера
    document.getElementById('entropy-display').innerText = data.entropy;

    // 2. Обновляем финальность
    document.getElementById('ms-value').innerText = data.finality.toFixed(3);
    document.getElementById('finality-bar').style.width = (data.finality * 100) + "%";

    // 3. Реакция на транзакцию
    if (data.is_processing) {
        triggerTransactionPulse();
    }
};

socket.onerror = () => {
    // Если сервер не запущен, включаем автономную имитацию
    runOfflineDemo();
};

socket.onclose = () => {
    runOfflineDemo();
};

// --- РЕЖИМ АВТОНОМНОЙ ДЕМОНСТРАЦИИ (Если сервер оффлайн) ---
function runOfflineDemo() {
    if (window.demoActive) return;
    window.demoActive = true;
    
    document.getElementById('terminal-text').innerText = "OFFLINE MODE: RUNNING LOCAL SIMULATION...";
    
    setInterval(() => {
        // Имитируем байты энтропии
        const fakeEntropy = Math.random().toString(16).slice(2, 10).toUpperCase();
        document.getElementById('entropy-display').innerText = fakeEntropy;

        // Имитируем финальность
        const fakeFinality = (0.7 + Math.random() * 0.25).toFixed(3);
        document.getElementById('ms-value').innerText = fakeFinality;
        document.getElementById('finality-bar').style.width = (fakeFinality * 100) + "%";

        // Случайные импульсы
        if (Math.random() > 0.95) triggerTransactionPulse();
    }, 200);
}

// --- ИНТЕРФЕЙСНЫЕ СОБЫТИЯ ---

// Panic Switch
document.getElementById('panic-switch').addEventListener('click', () => {
    const confirmed = confirm("WARNING: ACTIVATE EMERGENCY PHYSICAL ZEROING?");
    if (confirmed) {
        targetRotationSpeed = 1.0; // Гипер-вращение
        document.body.style.filter = "invert(1) saturate(5) hue-rotate(180deg)";
        document.getElementById('terminal-text').innerText = "KEY DISSOLUTION IN PROGRESS... SYSTEM PURGED.";
        setTimeout(() => location.reload(), 2500);
    }
});

// Адаптация под размер окна
window.addEventListener('resize', () => {
    camera.aspect = container.clientWidth / container.clientHeight;
    camera.updateProjectionMatrix();
    renderer.setSize(container.clientWidth, container.clientHeight);
});

