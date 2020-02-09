## Generate 3D cellular automata

### Looking for new rules

Generate 16 random 5 states cellular automata after (64 iterations).

```
$ mkdir -p explore
$ bin/search -output explore -state 5 -fillseed 0.4 -iteration 64 -nb 16
```

### Growing an interesting rule

#### Output each iterations (for animation)

```
$ mkdir -p explore
$ bin/grow -output explore -input output-XXXXXXX.obj -iteration 128 -worldsize 300
```

#### Output only last iteration

```
$ mkdir -p explore
$ bin/grow -output explore -input output-XXXXXXX.obj -lastonly -iteration 128 -worldsize 300
```

### Demos

![demo](https://i.gyazo.com/48f9e3d10fca5472f4971fc672896717.png)

![demo](https://i.gyazo.com/99056ee4d79cba57a6073287c6187c8a.png)

![demo](https://media.giphy.com/media/L3L6fgN9HO6mj5AokY/giphy.gif)

![demo](https://media.giphy.com/media/YOBFy0WPAfrueMDYvz/giphy.gif)

More examples :
- https://www.youtube.com/watch?v=W0iA-pO_XaU
- https://www.behance.net/gallery/90505477/Cellular-Automata-3D