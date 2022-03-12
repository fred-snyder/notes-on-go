# Go naming conventions

1. Name variables using only abbreviations in small scopes
2. Use the full words in larger scopes
3. Don't repeat words
    - GOOD: player.Score
    - BAD: player.PlayerScore
3. Use camelCase (and first Capital for exported functions and types)
4. Use UPPERCASE - only - for acronyms
5. Don't use underscores (operators are not allowed by the compiler)
6. Names should not contain an indication of the package. 
7. A method or function which returns an object is named as a noun (so Age, not GetAge). To change an object, use SetName (so SetAge).
8. Name constant identifiers with all uppercase letters (const PI = 3.15)
