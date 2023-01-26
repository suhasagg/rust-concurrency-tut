use rayon::prelude::*;

struct ListNode {
    val: i32,
    next: Option<Box<ListNode>>,
}

impl ListNode {
    fn new(val: i32) -> Self {
        ListNode { val, next: None }
    }
}

fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
    let mut lists = lists.into_par_iter().filter_map(|x| x).collect::<Vec<_>>();
    if lists.is_empty() {
        return None;
    }
    while lists.len() > 1 {
        let mut next = vec![];
        for (l1, l2) in lists.chunks(2).into_par_iter().flat_map(|x| x.chunks(2)) {
            next.push(merge_two_lists(l1.pop().unwrap(), l2.pop().unwrap()));
        }
        lists = next;
    }
