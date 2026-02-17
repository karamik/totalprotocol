/**
 * TOTAL Protocol: Sentinel Core v.8.0
 * Module: Entropy Harvester (TRNG based on Ring Oscillators)
 * * DESCRIPTION:
 * This module generates raw entropy by harvesting jitter from asynchronous 
 * ring oscillator chains. This is the physical root of trust for QRNG.
 */

module sentinel_entropy_harvester #(
    parameter int CHAIN_LENGTH = 5,
    parameter int SEED_WIDTH = 256
)(
    input  logic clk,
    input  logic rst_n,
    output logic [SEED_WIDTH-1:0] physical_entropy,
    output logic ready
);

    logic [CHAIN_LENGTH-1:0] ro_feedback;
    logic [SEED_WIDTH-1:0] shift_reg;
    int bit_counter;

    // Ring Oscillator Chain: Generating physical noise
    genvar i;
    generate
        for (i = 0; i < CHAIN_LENGTH; i++) begin : gen_ro
            if (i == 0)
                assign ro_feedback[i] = !(ro_feedback[CHAIN_LENGTH-1]);
            else
                assign ro_feedback[i] = !(ro_feedback[i-1]);
        end
    generate

    // Sampling jitter into the entropy pool
    always_ff @(posedge clk or negedge rst_n) begin
        if (!rst_n) begin
            shift_reg <= '0;
            bit_counter <= 0;
            ready <= 1'b0;
        end else begin
            shift_reg <= {shift_reg[SEED_WIDTH-2:0], ro_feedback[0]};
            if (bit_counter < SEED_WIDTH) begin
                bit_counter <= bit_counter + 1;
                ready <= 1'b0;
            end else begin
                physical_entropy <= shift_reg;
                ready <= 1'b1;
            end
        end
    end

endmodule
