import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv("data.csv")

fig = plt.figure()
ax = fig.gca()
ax.set_yscale("log")
ax.set_ylabel("wall time [seconds]")

ax.plot(range(df.shape[0]), df["mean_9"]/df["mean_9"].iat[0])
ax.plot(range(df.shape[0]), df["mean_8"]/df["mean_8"].iat[0])
ax.plot(range(df.shape[0]), df["mean_7"]/df["mean_7"].iat[0])
ax.plot(range(df.shape[0]), df["mean_6"]/df["mean_6"].iat[0])

fig.tight_layout()
fig.savefig("main.svg")
