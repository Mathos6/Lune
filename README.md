# Lune Glasses – MVP WebRTC (Simulation PC)

## Objectif
Simuler des lunettes de réalité augmentée légères connectées à un serveur distant,
en utilisant WebRTC pour le streaming vidéo temps réel et la transmission des inputs.

Aucun hardware spécifique n’est utilisé :
- le serveur tourne sur PC
- les lunettes sont simulées via un navigateur web

## Architecture générale

Client (lunettes simulées) :
- Navigateur web
- Affichage du flux vidéo WebRTC
- Envoi des inputs (clavier / souris)

Serveur AR :
- Génération des frames (rendu simple)
- Encodage vidéo temps réel
- Envoi du flux via WebRTC
- Réception et traitement des inputs

## Technologies

- Serveur : Python + aiortc
- Client : HTML + JavaScript
- Protocole : WebRTC
- Transport : UDP (via WebRTC)

## MVP – Fonctionnalités minimales

1. Le serveur génère une image (statique ou animée)
2. L’image est streamée en temps réel au client
3. Le client affiche la vidéo plein écran
4. Le client envoie un input
5. Le serveur reçoit l’input et réagit

## Objectif technique clé

Mesurer la latence :
Input utilisateur → réaction visuelle à l’écran

## Pourquoi WebRTC

- Streaming vidéo temps réel
- Faible latence
- Gestion du jitter et de la perte de paquets
- Architecture identique à la future version embarquée ARM

## Évolution future

- Remplacer le navigateur par un client ARM
- Remplacer les inputs simulés par capteurs réels
- Déployer le serveur sur une machine distante
