import matplotlib.pyplot as plt
import matplotlib.ticker
import numpy as np
import pandas as pd

df = pd.read_csv("data.csv")

# Wall time.
fig = plt.figure(1)
ax = fig.gca()

ax.errorbar(
    df["no_cores"],
    df["mean"],
    yerr = df["mean"] * df ["stddev"] * 0.03,
    fmt = "x",
    capsize = 5,
)
ax.plot(
    np.linspace(0.8, 20),
    df["mean"].iat[0] / np.linspace(0.8, 20),
    color = "k",
    linestyle = "dashed",
)

ax.set_xlabel("number of goroutines")
ax.set_ylabel("wall time [seconds]")
ax.set_xscale("log")
ax.set_yscale("log")
ax.set_xlim(0.8, 20)
ax.set_ylim(1, 25)
ax.xaxis.set_major_formatter(matplotlib.ticker.FormatStrFormatter("%.0f"))
ax.yaxis.set_major_formatter(matplotlib.ticker.ScalarFormatter())

fig.tight_layout()
fig.savefig("main.svg")

# Number of cores times wall time.
fig = plt.figure(2)
ax = fig.gca()

ax.errorbar(
    df["no_cores"],
    df["no_cores"] * df["mean"],
    yerr = df["no_cores"] * df["mean"] * df ["stddev"] * 0.03,
    fmt = "x",
    capsize = 5,
)
ax.plot(
    np.linspace(0, 17),
    df["mean"].iat[0] * np.ones_like(np.linspace(0, 17)),
    color = "k",
    linestyle = "dashed",
)

ax.set_xlabel("number of goroutines")
ax.set_ylabel("number of goroutines x wall time [seconds]")
ax.set_xlim(0, 17)
ax.set_ylim(0, 60)
ax.xaxis.set_major_formatter(matplotlib.ticker.FormatStrFormatter("%.0f"))
ax.yaxis.set_major_formatter(matplotlib.ticker.ScalarFormatter())

fig.tight_layout()
fig.savefig("main_1.svg")
