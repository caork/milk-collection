# %%
import matplotlib.pyplot as plt
import pandas as pd

# %%
index = "7 -> 0 -> 1 -> 12 -> 15 -> 3 -> 11 -> 14 -> 18 -> 2 -> 13 -> 4 -> 17 -> 16 -> 5 -> 19 -> 6 -> 9 -> 10 -> 8 -> 20 -> 7"
places = list(map(int, index.split(" -> ")))
# %%
df = pd.read_csv("distance.csv", header=None)
lines = df.reindex(places)

fig, ax = plt.subplots()
ax.scatter(df[1], df[2])
for i, txt in enumerate(df[0]):
    ax.annotate(txt-1, (df[1][i], df[2][i]))
ax.plot(lines[:][1], lines[:][2])
plt.show()
