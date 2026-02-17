// --- CONFIGURATION ---
// Убедись, что эта ссылка открывается в браузере и выдает JSON
const VERCEL_API = 'https://totalprotocol.vercel.app/api/status';

// --- THREE.JS 3D SCENE SETUP ---
const container = document.getElementById('xray-container');
const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, container.clientWidth / container.clientHeight, 0.1, 1000);
const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
renderer.setSize(container.clientWidth, container.clientHeight);
container.appendChild(renderer.domElement);

// Геометрия ядра (Sentinel Core)
const geometry = new THREE.OctahedronGeometry(1.5, 0);
const material = new THREE.MeshBasicMaterial({ color: 0x00ff41, wireframe: true, transparent: true, opacity: 0.8 });
const core = new THREE.Mesh(geometry, material);
scene.add(core);

// Внешняя оболочка (Shield)
const outerGeo = new THREE.IcosahedronGeometry(2.2, 1);
const outerMat = new THREE.MeshBasicMaterial({ color: 0x00ff41, wireframe: true, transparent: true, opacity: 0.1 });
const shield = new THREE.Mesh(outerGeo, outerMat);
scene.add(shield);

camera.position.z = 5;

let targetRotationSpeed = 0.01;
let currentRotationSpeed = 0.01;

// Эффект вспышки при транзакции
function triggerTransactionPulse() {
    targetRotationSpeed = 0.15;
    material.opacity = 1.0;
    setTimeout(() => {
        targetRotationSpeed = 0.01;
        material.opacity = 0.8;
    }, 800);
}

// Цикл анимации
function animate() {
    requestAnimationFrame(animate);
    currentRotationSpeed += (targetRotationSpeed - currentRotationSpeed) * 0.1;
    core.rotation.y += currentRotationSpeed;
    core.rotation.x += currentRotationSpeed / 2;
    shield.rotation.y -= 0.002;
    renderer.render(scene, camera);
}
animate();

// --- API SYNC LOGIC (The Connection) ---

async function syncWithVercel() {
    try {
        console.log("Fetching updates from Vercel...");
        const response = await fetch(VERCEL_API);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();
        console.log("Data received:", data);

        // Обновляем текстовые поля в HTML
        const entropyEl = document.getElementById('entropy-display');
        const finalityEl = document.getElementById('ms-value');
        const terminalEl = document.getElementById('terminal-text');
        const barEl = document.getElementById('finality-bar');

        if (entropyEl) entropyEl.innerText = data.entropy;
        if (finalityEl) finalityEl.innerText = data.finality.toFixed(4);
        
        // Визуализация прогресс-бара
        if (barEl) {
            const finalityPercent = data.finality * 100;
            barEl.style.width = finalityPercent + "%";
        }

        // Если идет обработка (is_processing), запускаем импульс
        if (data.is_processing) {
            triggerTransactionPulse();
            if (terminalEl) terminalEl.innerText = "SENTINEL CORE: ACCELERATING...";
        } else {
            if (terminalEl) terminalEl.innerText = "SYSTEM STATUS: ACTIVE | NODE: VERCEL-CLOUD";
        }

    } catch (error) {
        console.error("Connection to Vercel failed:", error);
        const terminalEl = document.getElementById('terminal-text');
        if (terminalEl) terminalEl.innerText = "NODE CONNECTION LOST. RETRYING...";
    }
}

// Опрос сервера каждые 1000мс (1 секунда)
setInterval(syncWithVercel, 1000);

// Подстройка под размер окна
window.addEventListener('resize', () => {
    camera.aspect = container.clientWidth / container.clientHeight;
    camera.updateProjectionMatrix();
    renderer.setSize(container.clientWidth, container.clientHeight);
});
