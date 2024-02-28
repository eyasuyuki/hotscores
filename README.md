hotscores
====

Calc hotspot score from Subversion repository

# Hotspot score

$$Score = \sum_{i=0}^n {1 \over 1e^{-12ti+12}}$$

- n: bug fix times
- ti: bug fix timestamp between 0 and 1 (file create time=0, now=1)
