class Node:
	def __init__(self, id):
		self.id = id
		self.left = self.right = None

	def invert(self):
		if self.left is not None:
			self.left.invert()
		if self.right is not None:
			self.right.invert()
		self.left, self.right = self.right, self.left

	def __repr__(self):
		return f'({self.id}L: {self.left} {self.id}R: {self.right})'

root = Node(0)
root.right = Node(1)
root.left = Node(2)
root.right.left = Node(3)
root.left.left = Node(4)
root.left.right = Node(5)
print("original:")
print(root)
root.invert()
print("inverted:")
print(root)