using System;
using formula;

namespace main
{
    static class main
    {
        static void Main(string[] args)
        {
            string input1   = Environment.GetEnvironmentVariable("INPUT_TEXT");
            string input2   = Environment.GetEnvironmentVariable("INPUT_LIST");
            string input3   = Environment.GetEnvironmentVariable("INPUT_BOOLEAN");
            string input4   = Environment.GetEnvironmentVariable("INPUT_PASSWORD");

            new formula.Hello(input1, input2, input3, input4);
        }
}
}
