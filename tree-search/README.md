# Shallowest Duplicate Tree

## Algorithm

Based off on the assumption that the requested function is to find the first occurence of a duplicate node regardless of the level separation, the implemented algorithm is a Breadth-First Search to ensure the first duplicate encountered, if any, is the shallowest one.

## Considerations

Given that the chosen algorithm increases its time complexity the more levels are added to the subtree, an early return was added to ensure the search stops as soon a duplicate is found, reducing memory usage.
