import re

import pandas as pd

df = pd.DataFrame()

for exp_it in range(6, 9):
    means_it = []
    stddevs_it = []

    for opt_it in range(6):
        with open(f"../one-billion-row-challenge-golang/hans-XMG-EVO-M24/cpu_run{opt_it}_{exp_it}.stat") as f:
            mat = re.search(r"(\w+\.\w+) ± ([^% ]+)", f.read())
            assert mat is not None
            means_it.append(float(mat.group(1).replace("m", "e-3")))
            stddevs_it.append(int(mat.group(2)))

    df[f"mean_{exp_it}"] = means_it
    df[f"stddev_{exp_it}"] = stddevs_it

means_it = []

for opt_it in range(6):
    with open(f"../one-billion-row-challenge-golang/hans-XMG-EVO-M24/cpu_run{opt_it}_9.stat") as f:
        mat = re.search(r"(\w+\.\w+) ± ([^% ]+)", f.read())
        assert mat is not None
        means_it.append(float(mat.group(1)))

df[f"mean_9"] = means_it

df.to_csv("data.csv", index=False)
