module total_top (
    input wire clk,
    input wire rst_n,
    
    // Входы от физических датчиков
    input wire [63:0] qrng_raw_stream,
    input wire [7:0] sensor_zone_1,
    input wire [7:0] sensor_zone_2,
    
    // Выходы управления
    output wire p_cooling_act,      // Сигнал на Пельтье
    output wire system_halt,        // Полная остановка (Sentinel)
    output wire security_alarm      // Флаг для софта (Oracle)
);

    wire entropy_alarm;
    wire thermal_shutdown;

    // 1. Инстанс сторожа энтропии
    entropy_watchdog watchdog_inst (
        .clk(clk),
        .reset_n(rst_n),
        .entropy_data(qrng_raw_stream),
        .alarm(entropy_alarm)
    );

    // 2. Инстанс термо-менеджера (v.7.0)
    thermal_manager thermal_inst (
        .clk(clk),
        .temp_sensor_a(sensor_zone_1),
        .temp_sensor_b(sensor_zone_2),
        .peltier_enable(p_cooling_act),
        .critical_shutdown(thermal_shutdown)
    );

    // Логика объединения тревог
    assign system_halt = thermal_shutdown;
    assign security_alarm = entropy_alarm | thermal_shutdown;

endmodule
