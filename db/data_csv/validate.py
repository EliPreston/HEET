import csv

with open("ENERGY_STAR_Certified_Light_Fixtures_-_Downlights_20250607.csv", newline='', encoding='utf-8') as f:
    reader = csv.reader(f)
    header = next(reader)
    for i, row in enumerate(reader, start=1):
        if len(row) != len(header):
            print(f"Row {i} has {len(row)} columns, expected {len(header)}: {row}")

print('done')