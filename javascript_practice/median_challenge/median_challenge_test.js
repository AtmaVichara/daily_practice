const ArrayMedian = require('./median_challenge');
const chai = require("chai");
const should = chai.should();

describe("Array Median Test", () => {
  describe("Constructor Test", () => {
    it("should have empty array and sum", () => {
      let medianTracker = new ArrayMedian(1000);

      medianTracker.sum.should.eq(0)
      medianTracker.limit.should.eq(1000)
    })
  })
  describe("Methods", () => {
    describe("#append", () => {
      it("append random number to array, update counter & sum", () =>{
        let medianTracker = new ArrayMedian(1);
        medianTracker.append()

        medianTracker.array.should.have.lengthOf(1)
        medianTracker.array[0].should.be.a('number')
        medianTracker.sum.should.eq(medianTracker.array[0])
        medianTracker.count.should.eq(1)
      })
    })

    describe("#findMedian", () => {
      it("finds median of array", () => {
        let medianTracker = new ArrayMedian(3)
        medianTracker.append()
        medianTracker.append()
        medianTracker.append()

        let median = medianTracker.findMedian()

        median.should.eq(medianTracker.sum / medianTracker.count)
      })
    })

    describe("#run", () => {
      it("returns array of medians", () => {
        let mt = new ArrayMedian(3)
        let medians = mt.runner()

        medians[0].should.eq(mt.array[0] / 1)
        medians[1].should.eq((mt.array[0] += mt.array[1]) / 2)
        medians[2].should.eq(mt.sum / mt.count)
      })
    })
  })
})
