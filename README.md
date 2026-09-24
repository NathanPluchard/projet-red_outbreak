# OUTBREAK

Un jeu de rôle/survie textuel, jouable entièrement dans un terminal, écrit en Go.

## Contexte

OUTBREAK a été développé dans le cadre d'un projet d'apprentissage du langage Go. L'objectif était de dépasser le simple exercice de syntaxe pour construire un vrai jeu jouable, avec une boucle de gameplay complète : création de personnage, exploration, combat au tour par tour, gestion de ressources et progression du personnage — tout en repoussant les limites de ce qu'on peut afficher proprement dans un terminal (couleurs, dégradés, alignement).

## Fonctionnalités

- **Création de personnage** avec 4 classes jouables (Survivant, Punk, Médecin de fortune, Sauveur), chacune avec ses propres statistiques et spécialités.
- **Refuge à gérer** : améliorations du générateur, du stockage, des défenses et de l'atelier, cycle jour/nuit avec risque de raid nocturne.
- **Expéditions** dans 5 zones (Centre-ville, Hôpital abandonné, Usine, Base militaire, Laboratoire), chacune avec plusieurs vagues d'infectés et un combat de boss.
- **Combat au tour par tour** : tirer, frappe spéciale, soin, grenade, fuite.
- **Ressources à gérer** : ferraille, composants, médicaments, carburant, nourriture.
- **Survivants à recruter**, chacun avec un bonus passif pour le refuge.
- **Réputation** qui évolue avec les actions du joueur.
- **Économie** : marché noir, atelier de fabrication, arsenal (7 armes), forgeron/bricoleur.
- **Casino** ("la Zone Morte") avec un Blackjack complet (split, assurance, abandon, sabot à deux paquets) et une roulette.
- **Inventaire et objets** : potions de soin, potions de poison, sorts.
- **Journal de survie** : suivi des statistiques (infectés éliminés, boss vaincus, zones nettoyées, missions).
- **Interface terminal personnalisée** : cadres et barres de vie/mana en dégradé de couleurs vraies (24 bits), logo ASCII généré, alignement calculé caractère par caractère pour gérer proprement les emojis et les codes couleur.

## Prérequis d'installation

- [Go](https://go.dev/dl/) en version **1.21** ou supérieure.
- Un terminal supportant les couleurs ANSI 24 bits (la plupart des terminaux récents : Windows Terminal, iTerm2, GNOME Terminal, VS Code, etc.).

Vérifier sa version de Go installée :

```bash
go version
```

## Guide d'installation et de lancement

1. Récupérer le projet (cloner le dépôt ou extraire l'archive) puis se placer dans le dossier `src` :

   ```bash
   cd src
   ```

2. Lancer le jeu directement avec Go :

   ```bash
   go run .
   ```

   Ou bien compiler un exécutable puis le lancer :

   ```bash
   go build -o outbreak .
   ./outbreak        # ou outbreak.exe sous Windows
   ```

3. Suivre les instructions à l'écran : création du survivant, choix de la classe, puis navigation dans le menu principal.

## Organisation du code source

Le projet est un module Go (`module outbreak`) organisé en un package `main` et plusieurs packages internes, chacun responsable d'une partie du jeu :

```
src/
├── go.mod              module Go (outbreak, go 1.21)
├── main.go             point d'entrée, boucle de jeu principale, refuge,
│                        combat, marché, atelier, cycle jour/nuit
├── classes/
│   ├── classes.go       structure du personnage, classes jouables, fiche
│   └── ui.go            moteur d'affichage terminal (cadres, barres,
│                          couleurs 24 bits, logo ASCII)
├── casino/
│   └── casino.go        mini-jeux : Blackjack et roulette
├── monstre/
│   └── monstre.go        types de monstres, attaques, logique de combat
├── marchand/
│   └── marchand.go       marché noir (achat/vente de ressources)
├── forgeron/
│   └── forgeron.go       bricoleur / équipement par paliers
├── inventaire/
│   └── inventaire.go     gestion de l'inventaire du joueur
├── potion/
│   ├── potion.go
│   ├── vie.go            potions de soin
│   └── poison.go         potions de poison
├── sort/
│   └── sort.go           sorts utilisables par le joueur
└── expedition/
    └── expedition.go     module d'expédition alternatif (zones, vagues,
                            boss) — non branché au menu principal actuel
```

Chaque dossier correspond à un package Go indépendant, importé par `main.go` (à l'exception du package `expedition`, laissé de côté pour le moment).
