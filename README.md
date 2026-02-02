# Entropy Bridge 📡

**Acoustic Data Transmission over Air-Gapped Devices using Golang.**

Entropy Bridge — это акустический модем, который позволяет передавать данные между устройствами через звуковые волны (микрофон и динамики), не используя интернет, Bluetooth или Wi-Fi.

## 🔥 Особенности
* **FSK Модуляция:** Использует частоты **2100 Гц** (логический 0) и **3100 Гц** (логическая 1).
* **Anti-Noise Protection:** Специальный алгоритм "разогрева" обходит шумоподавление Windows.
* **Smart Filtering:** Приемник автоматически отсеивает шум и показывает только переданный текст.

## 🛠 Технологии
* **Go** (Golang)
* **PortAudio** (Low-level audio I/O)
* **Алгоритм Герцеля** (Goertzel algorithm) для распознавания частот.

## 🚀 Как запустить

### 1. Скачайте код
```bash
git clone [https://github.com/mordaHQ/entropy-bridge.git](https://github.com/mordaHQ/entropy-bridge.git)
cd entropy-bridge
