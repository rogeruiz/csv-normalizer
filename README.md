# CSV Normalizer

This is a cli-tool to "normalize" CSV files. There are a couple of requirements
that need to be met. The list below shows what requirements have or have not
been met.

- [x] UTF-8 character set
- [x] The column **Timestamp** is formatted in `RFC3339`.
- [x] The column **Timestamp** is converted from US/PST to US/EST.
- [x] The column **ZIP** is converted to being five-digits long with left-padded
  zeros.
- [x] The column **FullName** is converted to _UPPERCASE_.
- [x] The column **Address** is Unicode validated but kept as-is.
- [ ] The columns **FooDuration** and **BarDuration** are converted to seconds.
- [ ] The column **TotalDuration** is the sum of **FooDuration** and
  **BarDuration** columns.
- [x] The column **Notes** is free text but any invalid UTF-8 characters are
  converted to the `U+FFFD` Unicode Replacement Character.
- [ ] If `U+FFFD` makes invalid data, print a message to `STDERR` and drop the
  row from the output.

This project includes the `.csv` files from the original
`trussworks/truss-interview` repository to help facilitate testing. These files
are kept in the `data/` directory at the root of this project.

## Project tooling

This project leverages `Make` for its automation. To get started, you can run
the `help` target to see what targets there are.

```sh
>_ make help
```

### Building the project

To build the project, run the `build` target. This will create a `normalizer`
binary in a `tmp` directory.

```sh
>_ make build
```

### Running the project

After the previous command runs successfully, run the `normalizer` binary
directly.

```sh
>_ tmp/normalizer < data/sample.csv
```

The above command can also be invoked with the `run` target.

```sh
>_ make run
```
