# Nova - Over The Air Updates

<img width="400px" src="https://github.com/novaota/nova/blob/master/nova_logo.png?raw=true">

Eine Reihe von Komponenten, hauptsächlich in GO geschrieben um OTA Updates für Fahrzeuge zu erproben.
Diese sind im Rahmen einer Bachelorarbeit an der OTH Regensburg bei der PENTASYS AG in München entstanden.

## Vorraussetzungen
1. PostgreSQL Datenbank
2. MQTT Broker

## Services

### DeviceManagement

Eine REST API um Geräte (Fahrzeuge), Besitzer, Updates und UpdateTasks zu verwalten.

### Notification Service

Eine Komponente um Fahrzeuge über MQTT über Updates zu informieren.

### Status Receiver

Eine Komponente um Status-Updates von Fahrzeugen in einer Datenbank zu persistieren.

### CarService

Beispielimplementierung eines Services, der ein Auto simuliert.

### UpdateEngine

Führt beliebige Updates auf dem Auto auf.

### Nova UI

Eine Beispieloberfläche einer HeadUnit geschrieben in QT C++.
