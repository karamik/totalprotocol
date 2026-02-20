/*
 * TOTAL Protocol: Sentinel Core v.8.1
 * Module: NTT Butterfly Unit (Dilithium PQC Accelerator)
 * Purpose: Hardware acceleration for lattice-based polynomial multiplication.
 */

module ntt_butterfly (
    input wire clk,
    input wire rst_n,
    input wire [22:0] a,      // Коэффициент A (23-bit для q=8380417)
    input wire [22:0] b,      // Коэффициент B
    input wire [22:0] zeta,   // Корень из единицы (Twiddle factor)
    output reg [22:0] u,      // Результат: u = a + b * zeta
    output reg [22:0] v       // Результат: v = a - b * zeta
);

    parameter Q = 23'd8380417;

    reg [45:0] mul_res;
    wire [22:0] b_zeta;

    // Шаг 1: Умножение (в реальном дизайне нужен конвейер Montgomery Reduction)
    always @(posedge clk or negedge rst_n) begin
        if (!rst_n) begin
            mul_res <= 46'd0;
        end else begin
            mul_res <= b * zeta;
        end
    end

    // Упрощенная редукция по модулю Q
    assign b_zeta = mul_res % Q;

    // Шаг 2: Сложение и вычитание (Butterfly)
    always @(posedge clk or negedge rst_n) begin
        if (!rst_n) begin
            u <= 23'd0;
            v <= 23'd0;
        end else begin
            // Результат u = (a + b_zeta) mod Q
            u <= (a + b_zeta >= Q) ? (a + b_zeta - Q) : (a + b_zeta);
            
            // Результат v = (a - b_zeta) mod Q
            v <= (a < b_zeta) ? (a + Q - b_zeta) : (a - b_zeta);
        end
    end

endmodule
