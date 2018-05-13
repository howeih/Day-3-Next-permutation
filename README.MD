#Day 3: Next permutation
---
Given a totally ordered set, find the next permutation of the current configuration: FADE -> FAED -> FDAE -> FDEA -> ...
<br/>
Even though almost every library offers a similar function, this tool can quickly become very useful in various competitions when you need to speed your algorithm up.
<br/>
The algorithm itself is simple. Beginning from end, find the longest decreasing sequence 42531 and denote the preceding item as pivot 42531. Swap pivot with the smallest higher item 43521 in the sequence and revert the sequence 43125.