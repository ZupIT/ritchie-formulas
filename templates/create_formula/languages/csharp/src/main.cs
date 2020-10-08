using System;
using formula;

namespace main
{
    static class main
    {
        static void Main(string[] args)
        {
            string input1   = Environment.GetEnvironmentVariable("input_text");
            string input2   = Environment.GetEnvironmentVariable("input_list");
            string input3   = Environment.GetEnvironmentVariable("input_boolean");
            string input4   = Environment.GetEnvironmentVariable("input_password");

            new formula.Hello(input1, input2, input3, input4);
        }
}
}
