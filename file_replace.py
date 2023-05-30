def replace(path, search_exp, replace_exp):
    file = open(path, "r")
    replacement = ""
    for line in file:
        line = line.strip()
        changes = line.replace(search_exp, replace_exp)
        replacement = replacement + changes + "\n"
    file.close()
    fout = open(path, "w")
    fout.write(replacement)
    fout.close()


replace("pkg/dashboard/model_policy.go", 'MID *string `json:"MID,omitempty"`', 'MID *string `json:"_id,omitempty"`')
