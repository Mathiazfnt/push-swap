# push-swap

## Organisation du projet

```text
push-swap/
├── checker/                 # Programme qui rejoue les instructions et valide le tri
├── entiers/                 # Conversion des arguments en liste d'entiers
├── instructions/
│   ├── echanger/            # sa, sb, ss
│   ├── pousser/             # pa, pb
│   ├── rotation_haut/       # ra, rb, rr
│   └── rotation_bas/        # rra, rrb, rrr
├── piles/                   # Piles A et B communes aux instructions
├── push-swap/               # Programme principal push-swap
└── go.mod
```

## Rôle des packages

- `entiers` : `Parse` transforme les arguments de la ligne de commande en entiers et rejette les valeurs invalides ou dupliquées.
- `piles` : contient l'état partagé des piles `A` et `B`, ainsi que `GetA` et `GetB` pour en obtenir une copie.
- `instructions/echanger` : échange les deux premiers éléments d'une pile.
- `instructions/pousser` : déplace le sommet d'une pile vers l'autre.
- `instructions/rotation_haut` : place le sommet d'une pile en bas.
- `instructions/rotation_bas` : place le dernier élément d'une pile au sommet.
- `checker` : lit les arguments et les instructions, applique les opérations, puis vérifie que A est triée et B vide.
- `push-swap` : point d'entrée du programme qui devra produire les instructions de tri.

Les deux programmes sont distincts : chacun initialise ses propres piles au démarrage. Dans un même programme, toutes les opérations utilisent les mêmes variables `piles.A` et `piles.B`.

Lancer les tests avec :

```sh
go test ./...
```
