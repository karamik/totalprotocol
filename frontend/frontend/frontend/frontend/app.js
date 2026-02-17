
const VERCEL_API = 'https://totalprotocol.vercel.app/api/status';

async function syncWithVercel() {
    try {
        const response = await fetch(VERCEL_API);
        if (!response.ok) throw new Error('Network response was not ok');
        
        const data = await response.json();
        
        // Прямое обновление текста (самый надежный способ)
        const entropyEl = document.getElementById('entropy-display');
        const finalityEl = document.getElementById('ms-value');
        
        if (entropyEl) {
            entropyEl.innerText = data.entropy;
            console.log("Entropy updated:", data.entropy);
        }
        if (finalityEl) {
            finalityEl.innerText = data.finality.toFixed(4);
        }

    } catch (error) {
        console.error("Fetch error:", error);
        const terminalEl = document.getElementById('terminal-text');
        if (terminalEl) terminalEl.innerText = "RECONNECTING TO SENTINEL...";
    }
}

// Запуск цикла
setInterval(syncWithVercel, 2000); // Раз в 2 секунды для стабильности
syncWithVercel(); // И один раз сразу при загрузке
