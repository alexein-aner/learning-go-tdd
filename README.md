## Introduction

This is a collection of exercises to learn both the Go programming language and test-driven development (TDD). For more details, visit: [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests).

## The Cycle

1. Write a test.
2. Make the compiler pass.
3. Run the test, see that it fails, and check that the error message is meaningful.
4. Write enough code to make the test pass.
5. Refactor.

At first glance, this may seem tedious, but sticking to the feedback loop is important.

Not only does it ensure that you have relevant tests, but it also helps ensure you design good software by refactoring with the safety of tests.

Seeing the test fail is an important check because it allows you to see what the error message looks like. As a developer, it can be very challenging to work with a codebase when failing tests do not provide a clear idea of what the problem is.

By ensuring your tests are fast and setting up your tools so that running tests is simple, you can enter a state of flow when writing your code.

By not writing tests, you are committing to manually checking your code by running your software, which breaks your state of flow. You won't be saving yourself any time, especially in the long run.
