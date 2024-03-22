# Project Design Overview

## Project Description

This is a neurotech headset simulator built in Golang. It serves as a simulation for what a neurotech headset's output is supposed to look like. The simulator generates random EEG data and sends it to a server. The server then processes the data and sends it to a client. The client is disjoint and can be found as one of the other repositories in this organization.

Note that the logic here also serves as a reference for bulding other server side applications for passing neural data from the headset to the client.

## Design Decisions
