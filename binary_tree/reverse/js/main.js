class Node {
  constructor(id) {
    this.left = null
    this.right = null
    this.id = id
  }

  invert() {
    if (this.left) {
      this.left.invert()
    }
    if (this.right) {
      this.right.invert()
    }
    [this.left, this.right] = (() => [this.right, this.left])()
  }
}

const root = new Node(0)
root.right = new Node(1)
root.left = new Node(2)
root.right.left = new Node(3)
root.left.left = new Node(4)
root.left.right = new Node(5)
console.log("original:")
console.log(root)
root.invert()
console.log("inverted:")
console.log(root)