package testlist

import (
    "strings"
)

type elem struct {
    value string
    next *elem
}

type List struct {
    head *elem
}

func revert(x List) List{
    res := List{}
    for t := x.head; t != nil; t = t.next {
        res.head = &elem {
            value: t.value,
            next: res.head,
        }
    }
    return res
}

func MEmpty() List {
    lst := List{}
    return lst
}

func (x List) MConcat(y List) List {
    if x.head == nil {
        return y
    }
    if y.head == nil {
        return x
    }
    rx := revert(x)
    res := List{head: y.head}
    for t := rx.head; t != nil; t = t.next {
        res.head = &elem{
            value: t.value,
            next: res.head,
        }
    }
    return res
}

func (x List) Add(s string) List {
    return List{
        head: &elem {
            value: s,
            next: x.head,
        },
    }
}

func (x List) Remove(s string) List {
    // we need to define what to do in specific cases:
    // - if element is not in the list - return unchanged (no error)
    // - if element exists more than once - only first is removed (otherwise use Filter)
    pos := -1
    for t, i := x.head, 0; t != nil; t = t.next {
        if t.value == s {
            pos = i
            break
        }
        i++
    }
    if pos == -1 {
        return x
    }
    if pos == 0 {
        return List{head: x.head.next}
    }
    res := List{head: &elem{value: x.head.value}}
    curSrc := x.head
    curDst := res.head
    for i := 1; i < pos; i++ {
        curDst.next = &elem{value: curSrc.next.value}
        curSrc = curSrc.next
        curDst = curDst.next
    }
    curDst.next = curSrc.next.next
    return res
}

func (x List) Size() int {
    cnt := 0
    for t := x.head; t != nil; t = t.next {
        cnt++
    }
    return cnt
}

func (x List) ToSlice() []string {
    res := []string{}
    for t := x.head; t != nil; t = t.next {
        res = append(res, t.value)
    }
    return res
}

func FromSlice(arr []string) List {
    res := List{nil}
    for i := len(arr) - 1; i >= 0; i-- {
        res.head = &elem{
            value: arr[i],
            next: res.head,
        }
    }
    return res
}

func (x List) Find(p func(v string) (bool)) (string, bool) {
    for t := x.head; t != nil; t = t.next {
        if p(t.value) {
            return t.value, true
        }
    }
    return "", false
}

func (x List) Filter(p func(v string) (bool)) List {
    res := List{}
    for t := x.head; t != nil; t = t.next {
        if p(t.value) {
            res.head = &elem {
                value: t.value,
                next: res.head,
            }
        }
    }
    return revert(res)
}

func (x List) Map(f func(v string) (string)) List {
    res := List{}
    for t := x.head; t != nil; t = t.next {
        res.head = &elem{
            value: f(t.value),
            next: res.head,
        }
    }
    return revert(res)
}

func (x List) Reduce(f func(a, b string) (string)) string {
    if x.head == nil {
        // we may want different contract for empty lists, but
        // for test task it is not important and consistent with following case
        return ""
    }
    acc := x.head.value
    for t := x.head.next; t != nil; t = t.next {
        acc = f(acc, t.value)
    }
    return acc
}

func (x List) String() string {
    return "[" + strings.Join(x.ToSlice(), ", ") + "]"
}
