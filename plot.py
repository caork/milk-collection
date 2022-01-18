# %%
import matplotlib.pyplot as plt
import pandas as pd

# %%
index = """
1->0->9->12->15->13->17->4->2->18->14->3->11->7->20->8->10->6->19->5->16
"""
index = index.strip()
places = list(map(int, index.split("->")))

# %%
df = pd.read_csv("distance.csv", header=None)
lines = df.reindex(places)

fig, ax = plt.subplots()
ax.scatter(df[1], df[2])
# for i, txt in enumerate(df[0]):
#     ax.annotate(txt-1, (df[1][i], df[2][i]))
ax.plot(lines[:][1], lines[:][2])
plt.show()
