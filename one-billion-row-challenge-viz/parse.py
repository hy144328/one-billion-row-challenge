import re

import pandas as pd

df = pd.DataFrame()

means_it = []
stddevs_it = []

for exp_it in range(16):
    with open(f"../one-billion-row-challenge-golang-parallel/hans-XMG-EVO-M24/cpu_run_{exp_it+1}.stat") as f:
        mat = re.search(r"(\w+\.\w+) ± ([^% ]+)", f.read())
        assert mat is not None
        means_it.append(float(mat.group(1).replace("m", "e-3")))
        stddevs_it.append(int(mat.group(2)))

df["no_cores"] = [i + 1 for i in range(16)]
df[f"mean"] = means_it
df[f"stddev"] = stddevs_it

df.to_csv("data.csv", index=False)
