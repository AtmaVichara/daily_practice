const Node = require('./node')

class LinkedList {

  constructor() {
    this.head = null
  }

  append(data, currentNode = this.head) {
    if (this.head === null) {
      this.head = new Node(data)
    } else if (currentNode.next == null) {
      currentNode.next = new Node(data)
    } else {
      currentNode = currentNode.next
      this.append(data, currentNode)
    }
  }

  sum(total = 0, currentNode = this.head) {
    if (currentNode === null) {
      return total
    } else {
      total += currentNode.data
      currentNode = currentNode.next
    }
    return this.sum(total, currentNode)
  }

  count() {
    if (this.head === null) {
      return 0
    }
    let currentNode = this.head
    let count = 0
    while (currentNode !== null) {
      currentNode = currentNode.next
      count++
    }
    return count
  }


}

module.exports = LinkedList
