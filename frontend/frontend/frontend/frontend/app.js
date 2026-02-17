/**
 * TOTAL PROTOCOL | Sentinel Core v.8.1
 * Frontend Logic & API Synchronization
 */

// 1. КОНФИГУРАЦИЯ (Проверь этот URL в браузере!)
const VERCEL_API = 'https://karamik-totalprotocol.vercel.app/api/status';

// 2. ИНИЦИАЛИЗАЦИЯ 3D СЦЕНЫ (Three.js)
let scene, camera, renderer, core, shield;
const container = document.getElementById('xray-container');

if (container) {
    init3D();
}

function init3D() {
    scene = new THREE.Scene();
    camera = new THREE.PerspectiveCamera(75, container.clientWidth / container.clientHeight, 0.1, 1000);
    
    renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
    renderer.setSize(container.clientWidth, container.clientHeight);
    container.appendChild(renderer.domElement);

    // Геометрия Sentinel Core
    const geometry = new THREE.OctahedronGeometry(1.5, 0);
    const material = new THREE.MeshBasicMaterial({ 
        color: 0x00ff41, 
        wireframe: true, 
        transparent: true, 
        opacity: 0.8 
    });
    core = new THREE.Mesh(geometry, material);
    scene.add(core);

    // Внешний щит
    const outerGeo = new THREE.IcosahedronGeometry(2.2, 1);
    const outerMat = new THREE.MeshBasicMaterial({ 
        color: 0x00ff41, 
        wireframe: true, 
        transparent: true, 
        opacity: 0.1 
    });
    shield = new THREE.Mesh(outerGeo, outerMat);
    scene.add(shield);

    camera.position.z = 5;
    animate();
}

function animate() {
    requestAnimationFrame(animate);
    if (core) {
        core.rotation.y += 0.01;
        core.rotation.x += 0.005;
    }
    if (shield) {
        shield.rotation.y -= 0.002;
    }
    renderer.render(scene, camera);
}

// 3. СИНХРОНИЗАЦИЯ С BACKEND (Vercel Go)
async function syncSentinel() {
    try {
        const response = await fetch(VERCEL_API);
        if (!response.ok) throw new Error('Offline');
        
        const data = await response.json();

        // Обновляем данные на экране
        const entropyEl = document.getElementById('entropy-display');
        const finalityEl = document.getElementById('ms-value');
        const statusEl = document.getElementById('status-text');
        const terminalEl = document.getElementById('terminal-text');
        const barEl = document.getElementById('finality-bar');

        if (entropyEl) entropyEl.innerText = data.entropy;
        if (finalityEl) finalityEl.innerText = data.finality.toFixed(4);
        
        if (statusEl) {
            statusEl.innerText = 'ACTIVE';
            statusEl.style.color = '#00ff41';
        }

        if (terminalEl) {
            terminalEl.innerText = `[${data.timestamp}] CORE_SYNC_OK: ENTROPY_RECEIVED`;
        }

        if (barEl) {
            barEl.style.width = (data.finality * 100) + "%";
        }

    } catch (error) {
        console.error("Connection lost:", error);
        const statusEl = document.getElementById('status-text');
        const terminalEl = document.getElementById('terminal-text');
        
        if (statusEl) {
            statusEl.innerText = 'RECONNECTING';
            statusEl.style.color = '#ff3e3e';
        }
        if (terminalEl) terminalEl.innerText = "ATTEMPTING TO RE-ESTABLISH COUPLING...";
    }
}

// Запуск циклов
setInterval(syncSentinel, 1500); // Опрос каждые 1.5 сек
syncSentinel(); // Первый запуск сразу

// Подстройка при изменении размера экрана (важно для мобилок)
window.addEventListener('resize', () => {
    if (container && camera && renderer) {
        camera.aspect = container.clientWidth / container.clientHeight;
        camera.updateProjectionMatrix();
        renderer.setSize(container.clientWidth, container.clientHeight);
    }
});
