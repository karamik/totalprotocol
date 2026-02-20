module thermal_manager (
    input wire clk,                  // Тактовая частота
    input wire [7:0] temp_sensor_a,  // Датчик зоны А
    input wire [7:0] temp_sensor_b,  // Датчик зоны Б
    output reg peltier_enable,       // Включение охлаждения
    output reg critical_shutdown     // Экстренная остановка
);

    parameter TEMP_LIMIT = 8'd45;    // Порог 45°C
    parameter CRITICAL_LIMIT = 8'd60; // Порог 60°C

    always @(posedge clk) begin
        // Если обе зоны горячие — активируем Пельтье (v.7.0)
        if (temp_sensor_a > TEMP_LIMIT && temp_sensor_b > TEMP_LIMIT) begin
            peltier_enable <= 1'b1;
        end else begin
            peltier_enable <= 1'b0;
        end

        // Если температура критическая — рубим питание
        if (temp_sensor_a > CRITICAL_LIMIT || temp_sensor_b > CRITICAL_LIMIT) begin
            critical_shutdown <= 1'b1;
        end else begin
            critical_shutdown <= 1'b0;
        end
    end
endmodule
