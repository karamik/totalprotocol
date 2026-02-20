module entropy_watchdog (
    input wire clk,           // Тактовый сигнал
    input wire reset_n,       // Сброс (активный низкий)
    input wire [63:0] entropy_data, // Данные от QRNG
    output reg alarm          // Сигнал тревоги для Guard
);

    reg [63:0] last_data;
    reg [15:0] stuck_counter;

    always @(posedge clk or negedge reset_n) begin
        if (!reset_n) begin
            alarm <= 1'b0;
            stuck_counter <= 16'd0;
            last_data <= 64'd0;
        end else begin
            if (entropy_data == last_data) begin
                // Если данные не меняются, наращиваем счетчик
                if (stuck_counter >= 16'hFFFF) begin
                    alarm <= 1'b1; // ТРЕВОГА: Данные "залипли"
                end else begin
                    stuck_counter <= stuck_counter + 1;
                end
            end else begin
                // Данные изменились, всё в порядке
                stuck_counter <= 16'd0;
                last_data <= entropy_data;
                alarm <= 1'b0;
            end
        end
    end
endmodule
