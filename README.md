# tinyLLM

How can a machine that understands absolutely nothing about language learn relationships from text well enough to predict what comes next?

## Tokenizer V0

The first goal of tinyLLM is to establish the smallest reliable representation of text before introducing higher-level language concepts.

At this stage, tinyLLM does not know what a word, character, sentence, or meaning is.

It only receives encoded data.

### Design Choice: UTF-8 Bytes

Go strings are UTF-8 encoded, so Tokenizer V0 uses raw bytes as the model's atomic symbols.

Encoding converts a string into its underlying byte sequence:

```text
"cat" -> [99 97 116]
```

## Tokenizer V1

The second stage is for tinyLLM to establish the baseline statistical distribution of bytes across the corpus:

1. how many times each byte appears
2. what fraction of the corpus each byte represents
3. what the baseline probability of each byte is

### Baseline Probability

For each byte `b`, tinyLLM calculates:

`P(b) = count(b) / total bytes`

In plain English, the probability of a byte is the number of times that byte appears divided by the total number of bytes in the corpus.

For example, if `t` appears 500 times in a corpus containing 10,000 bytes:

`P(t) = 500 / 10000 = 0.05`

So `t` accounts for 5% of the corpus.

This gives tinyLLM a baseline for what is normal before it begins looking at relationships between neighboring bytes.
