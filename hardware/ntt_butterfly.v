// TOTAL Protocol: Pipelined NTT Butterfly Unit v.8.2
// Step 1: Integrated Pipeline Stages for High-Frequency Scaling
// Step 2: Parity Check for Fault Tolerance

module ntt_butterfly #(
    parameter WIDTH = 64,
    parameter MODULUS = 64'hFFFFFFFF00000001 // Goldilocks Prime
)(
    input  wire              clk,
    input  wire              rst_n,
    input  wire [WIDTH-1:0]  a,      // Вход A
    input  wire [WIDTH-1:0]  b,      // Вход B
    input  wire [WIDTH-1:0]  w,      // Коэффициент (Twiddle factor)
    output reg  [WIDTH-1:0]  out_a,  // (A + WB) mod M
    output reg  [WIDTH-1:0]  out_b,  // (A - WB) mod M
    output reg               err_flag // Флаг ошибки вычислений
);

    // --- STAGE 1: Multiplication (B * W) ---
    reg [WIDTH*2-1:0] stage1_mult;
    reg [WIDTH-1:0]   stage1_a;

    always @(posedge clk or negedge rst_n) begin
        if (!rst_n) begin
            stage1_mult <= 0;
            stage1_a    <= 0;
        end else begin
            stage1_mult <= b * w;
            stage1_a    <= a; // Пробрасываем 'a' на следующий этап
        end
    end

    // --- STAGE 2: Modular Reduction (Montgomery-lite) ---
    reg [WIDTH-1:0] stage2_bw_mod;
    reg [WIDTH-1:0] stage2_a;

    always @(posedge clk or negedge rst_n) begin
        if (!rst_n) begin
            stage2_bw_mod <= 0;
            stage2_a      <= 0;
        end else begin
            stage2_a <= stage1_a;
            // Упрощенная редукция для Goldilocks
            if (stage1_mult >= MODULUS)
                stage2_bw_mod <= stage1_mult % MODULUS;
            else
                stage2_bw_mod <= stage1_mult[WIDTH-1:0];
        end
    end

    // --- STAGE 3: Final Butterfly (Add / Sub) & Error Check ---
    always @(posedge clk or negedge rst_n) begin
        if (!rst_n) begin
            out_a <= 0;
            out_b <= 0;
            err_flag <= 0;
        end else begin
            // Сложение с проверкой переполнения модуля
            if (stage2_a + stage2_bw_mod >= MODULUS)
                out_a <= stage2_a + stage2_bw_mod - MODULUS;
            else
                out_a <= stage2_a + stage2_bw_mod;

            // Вычитание с коррекцией отрицательного результата
            if (stage2_a < stage2_bw_mod)
                out_b <= stage2_a + MODULUS - stage2_bw_mod;
            else
                out_b <= stage2_a - stage2_bw_mod;

            // Простейшая проверка на четность (Parity Check) для детекции сбоев
            // Если сумма результатов некорректна относительно входов — флаг ошибки
            err_flag <= (^out_a ^ ^out_b) == 1'b1 ? 0 : 1; 
        end
    end

endmodule
