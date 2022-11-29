<!-- ---
title: Rust lang
subtitle: rs
date: 2020-10-14
bigimg: [{src: "/primitives/img/unsplash-josiah-ingels.jpg", desc: "Path"}]
--- -->

# how to write test

```rs
#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_123() {
        // ...
    }
}
```

```rs
#[test]
fn test_123() {
    // ...
}

```