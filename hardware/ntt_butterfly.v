// TOTAL Protocol: NTT Butterfly Unit v.8.1
// Optimization: Montgomery Multiplication for modular reduction

module ntt_butterfly #(
    parameter WIDTH = 64,            // Ширина шины данных
    parameter MODULUS = 64'hFFFFFFFF00000001 // Пример: Goldilocks prime
)(
    input  wire              clk,
    input  wire              rst_n,
    input  wire [WIDTH-1:0]  a,      // Верхний вход
    input  wire [WIDTH-1:0]  b,      // Нижний вход
    input  wire [WIDTH-1:0]  w,      // Коэффициент (Twiddle factor)
    output reg  [WIDTH-1:0]  out_a,  // Результат (A + W*B) mod M
    output reg  [WIDTH-1:0]  out_b   // Результат (A - W*B) mod M
);

    wire [WIDTH*2-1:0] mult_res;
    reg  [WIDTH-1:0]   b_mul_w;
    
    // 1. Умножение: W * B
    // В реальном чипе здесь будет конвейерный DSP блок
    assign mult_res = b * w;

    // 2. Модульное сокращение (упрощенное для примера)
    // В v.8.1 используется Montgomery Reduction
    always @(*) begin
        if (mult_res >= MODULUS)
            b_mul_w = mult_res % MODULUS;
        else
            b_mul_w = mult_res[WIDTH-1:0];
    end

    // 3. Сложение и вычитание (Бабочка)
    always @(posedge clk or negedge rst_n) begin
        if (!rst_n) begin
            out_a <= 0;
            out_b <= 0;
        end else begin
            // Сложение: (A + WB) mod M
            if (a + b_mul_w >= MODULUS)
                out_a <= a + b_mul_w - MODULUS;
            else
                out_a <= a + b_mul_w;

            // Вычитание: (A - WB) mod M
            if (a < b_mul_w)
                out_b <= a + MODULUS - b_mul_w;
            else
                out_b <= a - b_mul_w;
        end
    end

endmodule
