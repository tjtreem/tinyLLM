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

