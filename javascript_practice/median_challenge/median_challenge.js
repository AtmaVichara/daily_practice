class ArrayMedian {

  constructor(limit) {
    this.array = new Array()
    this.sum = 0;
    this.count = 0;
    this.limit = limit;
  }

  runner(medians = new Array()) {
    if (this.count === this.limit) {
      return medians
    } else {
      this.append()
      medians.push(this.findMedian())
      return this.runner(medians)
    }
  }

  append() {
    let randomNumber = Math.floor(Math.random() * 100)
    this.array.push(randomNumber)
    this.count++
    this.sum += randomNumber
  }

  findMedian() {
    return this.sum / this.count
  }
}

module.exports = ArrayMedian
