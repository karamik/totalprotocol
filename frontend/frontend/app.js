
// --- CONFIGURATION ---
const VERCEL_API = 'https://totalprotocol.vercel.app/api/status';

// --- THREE.JS 3D SCENE (CORE) ---
const container = document.getElementById('xray-container');
const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, container.clientWidth / container.clientHeight, 0.1, 1000);
const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
renderer.setSize(container.clientWidth, container.clientHeight);
container.appendChild(renderer.domElement);

const geometry = new THREE.OctahedronGeometry(1.5, 0);
const material = new THREE.MeshBasicMaterial({ color: 0x00ff41, wireframe: true, transparent: true, opacity: 0.8 });
const core = new THREE.Mesh(geometry, material);
scene.add(core);

const outerGeo = new THREE.IcosahedronGeometry(2.2, 1);
const outerMat = new THREE.MeshBasicMaterial({ color: 0x00ff41, wireframe: true, transparent: true, opacity: 0.1 });
const shield = new THREE.Mesh(outerGeo, outerMat);
scene.add(shield);
camera.position.z = 5;

let targetRotationSpeed = 0.01;
let currentRotationSpeed = 0.01;

function triggerTransactionPulse() {
    targetRotationSpeed = 0.15;
    material.opacity = 1.0;
    setTimeout(() => {
        targetRotationSpeed = 0.01;
        material.opacity = 0.8;
    }, 800);
}

function animate() {
    requestAnimationFrame(animate);
    currentRotationSpeed += (targetRotationSpeed - currentRotationSpeed) * 0.1;
    core.rotation.y += currentRotationSpeed;
    core.rotation.x += currentRotationSpeed / 2;
    shield.rotation.y -= 0.002;
    renderer.render(scene, camera);
}
animate();

// --- NETWORK NODE CONNECTION (THE BRIDGE) ---

async function syncWithVercel() {
    try {
        const response = await fetch(VERCEL_API);
        const data = await response.json();

        // Передаем данные из API в интерфейс
        document.getElementById('entropy-display').innerText = data.entropy;
        document.getElementById('ms-value').innerText = data.finality.toFixed(4);
        
        // Визуализируем прогресс-бар финальности
        const finalityPercent = data.finality * 100;
        document.getElementById('finality-bar').style.width = finalityPercent + "%";

        // Если статус обработки (is_processing) true - ускоряем ядро
        if (data.is_processing) {
            triggerTransactionPulse();
            document.getElementById('terminal-text').innerText = "SENTINEL CORE: PROCESSING REAL-TIME ACCELERATION...";
        } else {
            document.getElementById('terminal-text').innerText = "SYSTEM STATUS: ACTIVE | NODE: VERCEL-CLOUD";
        }

    } catch (error) {
        console.error("Connection failed:", error);
        document.getElementById('terminal-text').innerText = "RECONNECTING TO CLOUD NODE...";
    }
}

// Опрашиваем твой Vercel каждые 800 миллисекунд
setInterval(syncWithVercel, 800);

// Адаптация экрана
window.addEventListener('resize', () => {
    camera.aspect = container.clientWidth / container.clientHeight;
    camera.updateProjectionMatrix();
    renderer.setSize(container.clientWidth, container.clientHeight);
});
