import matplotlib.pyplot as plt
import sympy

import approximations
import rational

# This script does the following:
# - Computes polynomial and rational approximations of a given function (e^x by default).
# - Computes (x,y) coordinates for every approximation given the same x coordinates.
# - Plots the results for rough comparison.
# - Computes the max error for every approximation given the same x coordinates.
# - Computes and plots max errors for every approximation with a varying number of parameters.
# The following are the resources used to write the script:
# https://xn--2-umb.com/22/approximation/
# https://sites.tufts.edu/atasissa/files/2019/09/remez.pdf
def main():

    ##########################
    # Configuration Parameters

    # start of the interval to calculate the approximation on
    x_start = 0
    # end of the interval to calculate the approximation on
    x_end = 1
    # number of paramters to use for the approximations.
    num_parameters = 13

    # number of paramters to use for plotting error deltas.
    num_parameters_errors = 30

    # number of (x,y) coordinates used to plot the resulting approximation.
    num_points_plot = 100000

    # function to approximate
    approximated_fn = lambda x: sympy.Pow(sympy.E, x)

    # flag controlling whether to plot each approximation.
    # Plots if true.
    shouldPlotApproximations = False

    # flag controlling whether to compute max error for each approximation
    # given the equally spaced x coordinates.
    # Computes if true.
    shouldComputeErrorDelta = False

    # flag controlling whether to plot errors over a range.
    # Currently, does so only for Chebyshev Rational Approximation.
    # Computes if true.
    shouldPlotErrorRange  = True

    # flag controlling whether to plot max error for every approximation
    # with a varying number of parameters. This is useful to find the most
    # optimal number of parameters to use for each kind of approximation.
    # Plots if true.
    shouldPlotMaxError = False

    # Equispaced x coordinates to be used for plotting every approximation.
    x_coordinates = approximations.linspace(x_start, x_end, num_points_plot)

    if shouldComputeErrorDelta or shouldPlotApproximations or shouldPlotErrorRange:
        ###############################################
        # Approximation With Given Number of Parameters
        y_eqispaced_poly, y_chebyshev_poly, y_chebyshev_rational, y_actual = approximations.approx_and_eval_all(approximated_fn, num_parameters, x_coordinates)

        ################
        # Compute Errors
        if shouldComputeErrorDelta:
            print(F"\n\nMax Error on [{x_start}, {x_end}]")
            print(F"{num_points_plot} coordinates equally spaced on the X axis")
            print(F"{num_parameters} parameters used\n\n")

            # Equispaced Polynomial Approximation
            max_error_equispaced_poly = approximations.compute_max_error(y_eqispaced_poly, y_actual)
            print(F"Equispaced Poly: {max_error_equispaced_poly.evalf(chop=1e-100)}")

            # Chebyshev Polynomial Approximation
            max_error_chebyshev_poly = approximations.compute_max_error(y_chebyshev_poly, y_actual)
            print(F"Chebyshev Poly: {max_error_chebyshev_poly.evalf(chop=1e-100)}")

            # Chebyshev Rational Approximation
            max_error_chebyshev_rational = approximations.compute_max_error(y_chebyshev_rational, y_actual)
            print(F"Chebyshev Rational: {max_error_chebyshev_rational.evalf(chop=1e-100)}")

        if shouldPlotErrorRange:
            chebyshev_rational_error_deltas = approximations.compute_error_range(y_chebyshev_rational, y_actual)

            plt.semilogy(x_coordinates, chebyshev_rational_error_deltas)

            plt.grid(True)
            plt.title(f"Chebyshev Rational e^x Errors on [{x_start}, {x_end}]. {num_parameters} params, {num_points_plot} points")
            plt.show()
            

        ###############################
        # Plot Every Approximation Kind
        if shouldPlotApproximations:
            # Equispaced Polynomial Approximation
            plt.plot(x_coordinates, y_eqispaced_poly, label="Equispaced Poly")

            # Chebyshev Polynomial Approximation
            plt.plot(x_coordinates, y_chebyshev_poly, label="Chebyshev Poly")

            # Chebyshev Rational Approximation
            plt.plot(x_coordinates, y_chebyshev_rational, label="Chebyshev Rational")

            # Actual With Large Number of Coordinates (evenly spaced on the X-axis)
            plt.plot(x_coordinates, y_actual, label=F"Actual")

            plt.legend(loc="upper left")
            plt.grid(True)
            plt.title(f"Appproximation of e^x on [{x_start}, {x_end}] with {num_parameters} parameters")
            plt.show()

    #####################################################
    # Calculate Errors Given Varying Number of Parameters
    if shouldPlotMaxError:
        x_axis = []

        deltas_eqispaced_poly = []
        deltas_chebyshev_poly = []
        deltas_chebyshev_rational = []

        ################
        # Compute Deltas
        # The deltas are taken from actual function values for different number of parameters
        # This is needed to find the most optimal number of parameters to use.
        for num_parameters_current in range(1, num_parameters_errors + 1):
            x_axis.append(int(num_parameters_current))
            y_eqispaced_poly, y_chebyshev_poly, y_chebyshev_rational, y_actual = approximations.approx_and_eval_all(approximated_fn, num_parameters_current, x_coordinates)

            deltas_eqispaced_poly.append(approximations.compute_max_error(y_eqispaced_poly, y_actual))
            deltas_chebyshev_poly.append(approximations.compute_max_error(y_chebyshev_poly, y_actual))
            deltas_chebyshev_rational.append(approximations.compute_max_error(y_chebyshev_rational, y_actual))

        ##################
        # Plot the results

        # Equispaced Polynomial Approximation
        plt.semilogy(x_axis, deltas_eqispaced_poly, label="Equispaced Poly")

        # Chebyshev Polynomial Approximation
        plt.semilogy(x_axis, deltas_chebyshev_poly, label="Chebyshev Poly")

        # Chebyshev Rational Approximation
        plt.semilogy(x_axis, deltas_chebyshev_rational, label="Chebyshev Rational")

        plt.legend(loc="upper left")
        plt.grid(True)
        plt.title(f"Approximation Errors on [{x_start}, {x_end}]")
        plt.gca().invert_yaxis()
        plt.xlabel('Number of Parameters')
        plt.ylabel(F"-log_10{{ max | f'(x) - f(x) | }} where x is in [{x_start}, {x_end}]")
        plt.show()

# This script isolates the 13-parameter Chebyshev Rational approximation of e^x
# We are planning to use it in production. Therefore, we need to peform coefficient
# truncations to 36 decimal points (the max osmomath supported precision).
def exponent_approximation_choice():

    ##########################
    # Configuration Parameters

    # start of the interval to calculate the approximation on
    x_start = 0
    # end of the interval to calculate the approximation on
    x_end = 1
    # number of paramters to use for the approximations.
    num_parameters = 13

    # number of paramters to use for plotting error deltas.
    num_parameters_errors = 30

    # number of (x,y) coordinates used to plot the resulting approximation.
    num_points_plot = 100000

    # fixed point precision used in Osmosis `osmomath` package.
    osmomath_precision = 36

    # function to approximate
    approximated_fn = lambda x: sympy.Pow(sympy.E, x)

    # Equispaced x coordinates to be used for plotting every approximation.
    x_coordinates = approximations.linspace(x_start, x_end, num_points_plot)
    x_coordinates = [sympy.Float(sympy.N(coef, osmomath_precision + 1), osmomath_precision + 1) for coef in x_coordinates]

    # Chebyshev Rational Approximation to get the coefficients.
    coef_numerator, coef_denominator = approximations.chebyshev_rational_approx(approximated_fn, x_start, x_end, num_parameters)

    # Truncate the coefficients to osmomath precision.
    coef_numerator = [sympy.Float(sympy.N(coef, osmomath_precision + 1), osmomath_precision + 1) for coef in coef_numerator]
    coef_denominator = [sympy.Float(sympy.N(coef, osmomath_precision + 1), osmomath_precision + 1) for coef in coef_denominator]

    # Evaluate approximation.
    y_chebyshev_rational = rational.evaluate(x_coordinates, coef_numerator, coef_denominator)

    # Compute Actual Values
    y_actual = approximations.get_y_actual(approximated_fn, x_coordinates)

    chebyshev_rational_error_deltas = approximations.compute_error_range(y_chebyshev_rational, y_actual)

    plt.semilogy(x_coordinates, chebyshev_rational_error_deltas)

    plt.grid(True)
    plt.title(f"Chebyshev Rational e^x Errors on [{x_start}, {x_end}]. {num_parameters} params, {num_points_plot} points")
    plt.show()

if __name__ == "__main__":
    #main()
    exponent_approximation_choice()
