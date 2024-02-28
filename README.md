hotscores
====

Calc hotspot score from Subversion repository

# Hotspot score

$$Score = \sum_{i=0}^n {1 \over 1+\exp{-12ti+12}}$$

- n: bug fix times
- ti: bug fix timestamp between 0 to 1 (file create time=0, now=1)
