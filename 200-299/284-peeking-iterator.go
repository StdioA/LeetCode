/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter        *Iterator
	nextPointer *int
}

func Constructor(iter *Iterator) *PeekingIterator {
	pi := &PeekingIterator{
		iter: iter,
	}
	pi.flush()
	return pi
}

func (this *PeekingIterator) flush() {
	if this.iter.hasNext() {
		next := this.iter.next()
		this.nextPointer = &next
	} else {
		this.nextPointer = nil
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.nextPointer != nil
}

func (this *PeekingIterator) next() int {
	defer this.flush()
	return this.peek()
}

func (this *PeekingIterator) peek() int {
	if this.nextPointer != nil {
		return *this.nextPointer
	}
	return 0
}