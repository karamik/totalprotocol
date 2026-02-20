// TOTAL Protocol: Entropy Watchdog v.8.1
// Мониторинг качества квантового шума в реальном времени.
// Защита от "залипания" (stuck-at faults) и статистических перекосов.

module entropy_watchdog #(
    parameter THRESHOLD = 8,      // Макс. кол-во одинаковых бит подряд
    parameter WINDOW_SIZE = 128   // Размер окна для проверки баланса 0 и 1
)(
    input  wire       clk,
    input  wire       rst_n,
    input  wire       entropy_bit,    // Входящий бит от QRNG
    output reg        entropy_valid,  // Сигнал: рандом качественный
    output reg [7:0]  error_code      // Код ошибки для Oracle Shield
);

    reg [3:0] consecutive_count;
    reg       last_bit;
    reg [7:0] ones_density; // Подсчет единиц в окне 128 бит
    reg [6:0] bit_counter;

    always @(posedge clk or negedge rst_n) begin
        if (!rst_n) begin
            consecutive_count <= 0;
            last_bit          <= 0;
            entropy_valid     <= 1;
            error_code        <= 8'h00;
            ones_density      <= 0;
            bit_counter       <= 0;
        end else begin
            // 1. Проверка на "залипание" (Stuck-at Fault)
            if (entropy_bit == last_bit) begin
                if (consecutive_count >= THRESHOLD) begin
                    entropy_valid <= 0;
                    error_code    <= 8'h01; // Ошибка: Слишком длинная серия
                end else begin
                    consecutive_count <= consecutive_count + 1;
                end
            end else begin
                consecutive_count <= 0;
                last_bit          <= entropy_bit;
            end

            // 2. Проверка плотности (Баланс 0 и 1)
            bit_counter <= bit_counter + 1;
            if (entropy_bit) ones_density <= ones_density + 1;

            if (bit_counter == WINDOW_SIZE - 1) begin
                // Если единиц меньше 25% или больше 75% — рандом "грязный"
                if (ones_density < (WINDOW_SIZE/4) || ones_density > (3*WINDOW_SIZE/4)) begin
                    entropy_valid <= 0;
                    error_code    <= 8'h02; // Ошибка: Статистический перекос
                end else if (error_code != 8'h01) begin
                    entropy_valid <= 1;
                    error_code    <= 8'h00;
                end
                ones_density <= 0;
            end
        end
    end

endmodule
