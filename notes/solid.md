
## Single Responsibility Principle (SRP)

- A class should have one primary responsibility and as a result should have one reason to change that is related to the primary responsibilities.
- `Separation Of Concerns` - Make sure that one package does one job only. For example, if you have a Journal Class, it should only be used to handle the journal.
- In case we want to add persistence (ie save journals to a file), you should create a separate package which can be used to save all kinds of objects.
- So later in the project, when you want to persist something again, you can use that package which would make further development better

## Open-Closed Principle (OCP)

- A class should be open for extension but closed for modification.
- `Specification Pattern` - Enterprise