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
git clone https://github.com/mordaHQ/entropy-bridge.git https://github.com/mordaHQ/entropy-bridge.git
cd entropy-bridge
```
## 2. Подготовка (Важно для Windows!)

Программа требует библиотеку PortAudio.

  **Скачайте файл portaudio_x64.dll.**

  **Положите его рядом с файлом main.go (в папки cmd/sender и cmd/receiver) перед запуском.**

## 3. Запуск Приемника (Receiver)

На принимающем компьютере:
```bash
go run cmd/receiver/main.go
```
## 4. Запуск Передатчика (Sender)

На передающем компьютере:
```bash
go run cmd/sender/main.go
```

Created by mordaHQ
