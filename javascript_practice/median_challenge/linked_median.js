const LinkedList = require("./linked_list")

class LinkedMedian {

  constructor(limit) {
    this.limit = limit
    this.ll = new LinkedList()
  }

  createMedianLinkedList() {
    let ml = new LinkedList()
    while (this.ll.count() !== this.limit) {
      let randomNumber = Math.floor(Math.random() * 100)
      this.ll.append(randomNumber)
      ml.append(this.median())
    }
    return ml
  }

  median() {
    if (this.ll.count() === 0) {
      return (this.ll.sum())
    } else {
      return (this.ll.sum() / this.ll.count())
    }
  }

}


module.exports = LinkedMedian
