const LinkedMedian = require('./linked_median');
const chai = require("chai");
const should = chai.should();
const expect = chai.expect

describe("Linked Median Test", () => {
  describe("Constructor Test", () => {
    it("should have empty array and sum", () => {
      let medianTracker = new LinkedMedian(1000);

      expect(medianTracker.ll.head).to.be.null
      medianTracker.limit.should.eq(1000)
    })
  })

  describe("#createMedianLinkedList", () => {
    it("should return new linked list of median values", () => {
      let mt = new LinkedMedian(3);
      let medianLL = mt.createMedianLinkedList()

      medianLL.count().should.eq(3)
      medianLL.head.data.should.eq(mt.ll.head.data)
      medianLL.head.next.data.should.eq((mt.ll.head.data + mt.ll.head.next.data) / 2)
      medianLL.head.next.next.data.should.eq(mt.ll.sum() / mt.ll.count())
    })
  })
})
