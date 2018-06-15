const LinkedList = require('./linked_list');
const chai = require("chai");
const should = chai.should();
const expect = chai.expect;

describe("Array Median Test", () => {
  describe("Constructor Test", () => {
    it("creates new linked list", () => {
      let ll = new LinkedList()

      expect(ll.head).to.be.null
    })
  })

  describe("#append", () => {
    it("appends node to list", () => {
      let ll = new LinkedList()
      ll.append(5)
      ll.append(4)
      ll.append(3)

      ll.head.data.should.eq(5)
      ll.head.next.data.should.eq(4)
      ll.head.next.next.data.should.eq(3)
    })
  })
  describe("#sum", () => {
    it("appends node to list", () => {
      let ll = new LinkedList()
      ll.append(5)
      ll.append(4)
      ll.append(3)
      let sum = ll.sum()

      sum.should.eq(12)
    })
  })
  describe("#count", () => {
    it("appends node to list", () => {
      let ll = new LinkedList()
      ll.append(5)
      ll.append(4)
      ll.append(3)
      let sum = ll.count()

      sum.should.eq(3)
    })
  })
})
