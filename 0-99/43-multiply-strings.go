func multiply(num1 string, num2 string) string {
    var(
        l1=len(num1)
        l2=len(num2)
        numBytes1=make([]byte, l1)
        numBytes2=make([]byte, l2)
        numBytes=make([]byte, l1 + l2)
    )
    for i: = l1 - 1
    i >= 0
    i-- {
        numBytes1[l1-i-1] = num1[i]-'0'
    }
    for i: = l2 - 1
    i >= 0
    i-- {
        numBytes2[l2-i-1] = num2[i]-'0'
    }

    for i, n1: = range numBytes1 {
        for j, n2: = range numBytes2 {
            numBytes[i+j] += n1 * n2
            if numBytes[i+j] >= 10 {
                numBytes[i+j+1] += numBytes[i+j] / 10
                numBytes[i+j] %= 10
            }
        }
    }
    var(
        valid bool
        buffer bytes.Buffer
    )
    for i: = len(numBytes) - 1
    i >= 0
    i-- {
        if numBytes[i] > 0 {
            valid = true
        }
        if valid {
            buffer.WriteByte(numBytes[i] + '0')
        }
    }
    if valid {
        return buffer.String()
    }
    return "0"
}
