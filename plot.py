# %%
import matplotlib.pyplot as plt
import pandas as pd

# %%
index = "10 -> 9 -> 6 -> 19 -> 5 -> 0 -> 16 -> 1 -> 17 -> 4 -> 13 -> 2 -> 15 -> 18 -> 14 -> 3 -> 12 -> 11 -> 7 -> 20 -> 8 -> 10"
places = list(map(int, index.split(" -> ")))
places =  [3, 2, 1, 5, 9, 12, 18, 16, 13, 10, 8, 11, 14, 15, 17, 6, 7, 4, 3]
# %%
df = pd.read_csv("distance.csv", header=None)
lines = df.reindex(places)

fig, ax = plt.subplots()
ax.scatter(df[1], df[2])
for i, txt in enumerate(df[0]):
    ax.annotate(txt-1, (df[1][i], df[2][i]))
ax.plot(lines[:][1], lines[:][2])
plt.show()
