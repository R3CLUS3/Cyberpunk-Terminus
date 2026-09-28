# 🌆 Cyberpunk Terminus

> **SYS.NET // TERMINAL INTERFACE v1.0**  
> Un jeu de rôle et de tactique textuel dans l'univers de Night City, codé en **Go**.

---

## 📜 Présentation

**Cyberpunk Terminus** est un RPG textuel en ligne de commande développé en Go. Incarnez un personnage de Night City, choisissez votre origine, gérez votre équipement, infiltrez des complexes hautement sécurisés en vue tactique 2D et affrontez les plus grandes légendes de la ville.

### 🛠️ Fonctionnalités clés

- **Création de Personnage & Progression** : Choisissez parmi 3 origines (*Gosse des Rues*, *Netrunner*, *Corpo*), gagnez de l'expérience et montez jusqu'au niveau 50.
- **Raids Tactiques 2D** :
  - **Militech Raid** : Infiltration sur 3 niveaux avec carte ASCII interactif (`Z`, `Q`, `S`, `D`).
  - **Arasaka Tower Raid** : Infiltration sur 2 niveaux face aux ninjas et à Adam Smasher.
- **Arène des Légendes** : Duel 1v1 contre les icônes de Night City (*David Martinez*, *Lucy*, *Rebecca*, *V*, *Adam Smasher*).
- **Atelier Cyberware & Marchand** : Achat d'équipements, d'implants et d'armures uniques (comme la *Veste de David Martinez* avec vérification biométrique).
- **Système de mort & Cyberpsychose** : Équipements expérimentaux à haut risque (*Cybersquelette Militech*) et réapparition sécurisée.
- **Interface Terminal Propre** : Rafraîchissement dynamique de l'écran avec codes de couleurs ANSI.

---

## 📋 Prérequis

Pour exécuter et jouer au jeu, vous devez disposer de **Go (Golang)** installé sur votre machine.

- **Go** : Version `1.18` ou supérieure.  
  👉 [Télécharger Go](https://go.dev/dl/)

---

## ⚙️ Installation

1. **Cloner le dépôt GitHub**
   ```bash
   git clone [https://github.com/R3CLUS3/Cyberpunk-Terminus.git](https://github.com/R3CLUS3/Cyberpunk-Terminus.git)
   cd Cyberpunk_Terminus
   go run main.go
   ou
   go run .