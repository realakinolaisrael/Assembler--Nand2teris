# Assembler- Nand2teris

This project is a simple Hack Assembler written in Go, based on the Nand2Tetris course. It translates .asm files written in Hack assembly language into binary .hack files that can run on the Hack computer platform.

🚀 What It Does
Reads a .asm file (written in Hack assembly language)

Removes comments and empty lines

Translates A-instructions (e.g., @2) into 16-bit binary code

Outputs the binary instructions into a .hack file

Has a basic placeholder for C-instructions (e.g., D=A, M=D+1)

🔧 Currently, this version only fully supports A-instructions. C-instruction support is a placeholder and can be expanded.

 ## 📁 File Structure
main.go – The main Go file containing all the assembler logic

example.asm – Sample Hack assembly program (create your own)

example.hack – Output file after assembly

🛠 How to Run
1. Create a .asm File
asm
Copy
Edit
// Simple addition
@2
D=A
@3
D=D+A
@0
M=D
Save it as example.asm.

2. Run the Assembler
bash
Copy
Edit
go run main.go example.asm
3. Output
A file called example.hack will be created, containing the translated binary instructions.

## ✍️ Beginner Tips
A-instructions start with @ and just contain numbers (e.g. @21)

C-instructions (like D=A, M=D+1) tell the Hack computer what to compute

Later, you can improve the assembler to fully handle C-instructions and symbolic labels

## 🧩 Next Steps
Add full support for C-instructions

Handle symbols (like variables and labels)

Add error checking for invalid syntax

## 📚 Resources
Nand2Tetris Course

Hack Machine Language Spec