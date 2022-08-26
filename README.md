## Generate 3D cellular automata

### Building

```
$ git clone https://github.com/tterrasson/crystal
$ cd crystal
$ export GO111MODULE="on"
$ go build -o bin/grow ./cmd/grow
$ go build -o bin/search ./cmd/search
```

### Looking for new rules

Generate 16 random 5 states cellular automata after 64 iterations.

```
$ bin/search -output explore -state 5 -fillseed 0.4 -iteration 64 -nb 16 -worldsize 300
```

### Growing an interesting rule

#### Output each iterations (for animation)

```
$ bin/grow -output explore -input output-XXXXXXX.obj -iteration 128 -worldsize 300
```

#### Output only last iteration

```
$ bin/grow -output explore -input output-XXXXXXX.obj -lastonly -iteration 128 -worldsize 300
```

### Demos - Rendered with [Blender](https://www.blender.org/) (Cycles)

![demo](https://i.gyazo.com/48f9e3d10fca5472f4971fc672896717.png)

![demo](https://i.gyazo.com/99056ee4d79cba57a6073287c6187c8a.png)

![demo](https://media.giphy.com/media/L3L6fgN9HO6mj5AokY/giphy.gif)

![demo](https://media.giphy.com/media/ZG0yw17IqGS9yZOGo0/giphy.gif)

![demo](https://media.giphy.com/media/eMmgOUTYCJRZzGiwbb/giphy.gif)

More examples :
- https://www.instagram.com/_cellular_automata_/
- https://www.youtube.com/watch?v=W0iA-pO_XaU
- https://www.behance.net/gallery/90505477/Cellular-Automata-3D
