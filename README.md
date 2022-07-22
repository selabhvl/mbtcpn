# MBT/CPN: Model-based Testing with Coloured Petri Nets

MBT/CPN is a tool for model-based testing with Coloured Petri Nets. The tool is used as a library on top of CPN Tools and requires installation of CPN Tools. The instructions below assumes that you are familiar with how to use CPN Tools.

## Download

MBT/CPN is distributed as a zip-file containing a set of Standard ML files that needs to be loaded into CPN Tools. Start by download the most recent version provided via this repository.

## Install

Unpack the distribution zip-file in some folder on your machine.

## Use

The examples folder in the zip-file contains an example CPN model `twophasecommit.cpn` that illustrates how to use MBT/CPN. The example can be found in the `examples/tpc` folder in the zip-file. To use the example you need to do the following steps:

1. Load the `twophasecommit.cpn` model into CPN Tools
2. Edit the declaration: `val mbtcpnlibpath =  "c:/work/git/mbtcpnlib/";` to match the folder in which you unpacked the MBT/CPN library
3. Enter the state space tool of CPN Tools
4. Open the MBT page/module and load the library by using ML Evaluate on the ML `use`-statements to load the MBT/CPN library.
5. Use `Execute.ss` and `Execute.sim <n>` to generate test cases using the state space-based approach and the simulation-based approach, respectively. 
6. Use `Execute.export ()` to export the test cases into XML files. These will be put in the output folder in the examples/tpc folder.

To use MBT/CPN in your own model, you need to define the `InEvent` and `OutEvent` colour sets to match the specific CPN model. The file `tcg.sml` in `examples/tpc` gives an example of how to implement the test case specification required by MBT/CPN to generate test cases.

## References

We have described the library in: [Rui Wang, Lars Michael Kristensen, Volker Stolz: **MBT/CPN: A Tool for Model-Based Software Testing of Distributed Systems Protocols Using Coloured Petri Nets**, in VECoS 2018, Lecture Notes in Computer Science 11181, Springer 2018](https://doi.org/10.1007/978-3-030-00359-3_7). Another paper that uses the same approach to generate test cases for a distributed system written in Go is: [Rui Wang, Lars Michael Kristensen, Hein Meling, Volker Stolz: **Automated test case generation for the Paxos single-decree protocol using a Coloured Petri Net model**, in J. Log. Algebraic Methods Program. 104, 2019](https://doi.org/10.1016/j.jlamp.2019.02.004).

If this library is useful to you, you may also be interested in our work of coverage analysis of CPNs: [Faustin Ahishakiye, José-Ignacio Requeno Jarabo, Lars Michael Kristensen, Volker Stolz: **Coverage Analysis of Net Inscriptions in Coloured Petri Net Models**, in VECoS 2020,  Lecture Notes in Computer Science 12519, Springer 2020](https://doi.org/10.1007/978-3-030-65955-4_6).
