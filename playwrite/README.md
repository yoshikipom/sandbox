# Commands
```bash
# open URL, operate the page -> generagte code
npx playwright codegen https://playwright.dev/

# execute all test (test folder is defined in playwrite.config.ts)
npx playwright test
npx playwright test tests/example.spec.ts

# parallel
npx playwright test --workers 4

# Show UI for tests
npx playwright test --ui

# report with trace
npx playwright test --trace on && npx playwright show-report
```

# Ref
- Official https://playwright.dev/docs/intro
- https://future-architect.github.io/articles/20230821a/
